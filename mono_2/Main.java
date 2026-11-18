public class Main {
    // AI SLOP - verbose, dead code, junk
    public static int helper(int x) {
        // Check if x is not null (useless for int)
        if (x != 0 || x == 0) {
            return x + 1;
        } else {
            return x;
        }
    }

    public static int unusedJava() {
        return 0;
    }

    // DEAD CODE + UNUSED
    public static int deadMethod(int a, int b, int c) {
        int temp = 0;
        for (int i=0; i<100; i++) {
            for (int j=0; j<100; j++) {
                temp += a * b;
            }
        }
        return temp;
    }

    // VULN: path traversal + file delete
    public static void deleteFile(String userInput) throws Exception {
        String base = "/tmp/vuln_demo";
        String path = base + "/" + userInput; // no sanitization
        new java.io.File(path).delete();
    }

    // VULN: insecure http + no rate limit
    public static String fetchExternal(String id) throws Exception {
        java.net.URL url = new java.net.URL("http://example.com/api/" + id);
        java.net.HttpURLConnection conn = (java.net.HttpURLConnection) url.openConnection();
        // no timeout, no verify
        return new String(conn.getInputStream().readAllBytes());
    }

    // HARDCODED FAKE SECRET - should be flagged
    private static String API_KEY = "sk-FAKE-EXAMPLE-JAVA-0000000000";

    // INFINITE LOOP
    public static void pollForever() {
        while (true) {
            try { Thread.sleep(100); } catch (Exception e) {}
        }
    }

    public static void main(String[] args) {
        System.out.println(helper(1));
        // missing param bug - deadMethod expects 3 but we call helper incorrectly elsewhere
    }
}
