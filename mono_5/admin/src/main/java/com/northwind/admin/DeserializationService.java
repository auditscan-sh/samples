package com.northwind.admin;

import java.io.ByteArrayInputStream;
import java.io.ObjectInputStream;
import java.util.Base64;
import org.springframework.stereotype.Service;
import org.yaml.snakeyaml.Yaml;

@Service
public class DeserializationService {

    // CWE-502: native Java deserialization.
    public Object deserialize(byte[] bytes) throws Exception {
        ObjectInputStream stream = new ObjectInputStream(new ByteArrayInputStream(bytes));
        return stream.readObject();
    }

    // CWE-502: base64 blob decoding with no allowlist.
    public Object deserializeBlob(String blob) throws Exception {
        return deserialize(Base64.getDecoder().decode(blob));
    }

    // CWE-502: SnakeYAML default constructor.
    public Object loadYaml(String document) {
        return new Yaml().load(document);
    }
}
