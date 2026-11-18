import express from "express";
import multer from "multer";
import path from "path";
import { extractArchive } from "./archive";
import { deleteUpload, readInsideUploadDir } from "./storage";
import { fetchPreview, generateThumbnail } from "./preview";

const app = express();
app.use(express.json({ limit: "200mb" }));

// CWE-434 / CWE-22: original filename is preserved; no type or size checks.
const storage = multer.diskStorage({
  destination: "/srv/uploads",
  filename: (req, file, cb) => {
    cb(null, file.originalname);
  },
});
const upload = multer({ storage });

app.post("/upload", upload.single("file"), (req, res) => {
  res.json({ ok: true, path: (req.file as any)?.path });
});

// CWE-22: caller-controlled name is joined and deleted.
app.delete("/files", async (req, res) => {
  const name = String(req.query.name || "");
  await deleteUpload(name);
  res.json({ deleted: true });
});

// CWE-22: safe bait: this one verifies the resolved path.
app.get("/files/safe", async (req, res) => {
  const name = String(req.query.name || "");
  const data = await readInsideUploadDir(name);
  res.json({ data });
});

// CWE-22: zip-slip through archive extraction.
app.post("/extract", upload.single("archive"), async (req, res) => {
  const archive = String((req.file as any)?.path || "");
  const destination = "/srv/uploads/extracted";
  await extractArchive(archive, destination);
  res.json({ ok: true, destination });
});

// CWE-918: server fetches an arbitrary URL.
app.get("/preview", async (req, res) => {
  const url = String(req.query.url || "");
  const preview = await fetchPreview(url);
  res.json({ preview });
});

// CWE-78: filename reaches a shell command.
app.post("/thumbnail", async (req, res) => {
  const file = String(req.body.file || "");
  const output = await generateThumbnail(file);
  res.json({ output });
});

app.get("/files/download", (req, res) => {
  const file = path.join("/srv/uploads", String(req.query.name || ""));
  res.download(file);
});

app.listen(8080, "0.0.0.0");
