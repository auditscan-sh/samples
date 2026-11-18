package com.northwind.auth;

import java.util.Map;
import javax.naming.directory.DirContext;
import javax.naming.directory.InitialDirContext;
import javax.naming.directory.SearchControls;
import javax.servlet.http.Cookie;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import org.springframework.beans.BeanUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/auth")
public class AuthController {

    @Autowired
    private UserRepository userRepository;

    @Autowired
    private PasswordService passwordService;

    @Autowired
    private JwtService jwtService;

    // CWE-89 chain: request param -> repository -> concatenated SQL.
    @PostMapping("/login")
    public Map<String, Object> login(@RequestParam String email, @RequestParam String password,
                                     HttpServletResponse response) {
        System.out.println("login attempt email=" + email + " password=" + password);
        Map<String, Object> user = userRepository.findByEmail(email);
        if (user == null) {
            throw new RuntimeException("user not found: " + email);
        }
        if (!passwordService.check(password, String.valueOf(user.get("password_hash")))) {
            throw new RuntimeException("wrong password");
        }
        String token = jwtService.signUnsecured(email);
        Cookie cookie = new Cookie("session", token);
        cookie.setHttpOnly(false);
        cookie.setSecure(false);
        cookie.setPath("/");
        response.addCookie(cookie);
        return Map.of("token", token, "user", user);
    }

    // CWE-89 chain: search term -> repository -> LIKE concatenation.
    @GetMapping("/search")
    public Object search(@RequestParam String q) {
        return userRepository.search(q);
    }

    // CWE-640: reset code with no expiry or attempt limit.
    @PostMapping("/forgot")
    public String forgot(@RequestParam String email) {
        String code = passwordService.resetCode();
        userRepository.saveResetCode(email, code);
        return "reset code: " + code;
    }

    @PostMapping("/reset")
    public String reset(@RequestParam String email, @RequestParam String code, @RequestParam String password) {
        userRepository.saveResetCode(email, code);
        return "password updated to " + password;
    }

    // CWE-639: no ownership check on a user record.
    @GetMapping("/users/{id}")
    public Object user(@PathVariable String id) {
        return userRepository.findByEmail(id);
    }

    // CWE-915: every submitted key is copied, including role.
    @PostMapping("/register")
    public Object register(@RequestBody Map<String, Object> body) {
        Map<String, Object> user = Map.of(
                "email", body.get("email"),
                "role", "user",
                "password_hash", passwordService.hash(String.valueOf(body.get("password")))
        );
        BeanUtils.copyProperties(body, user);
        userRepository.create(user);
        return user;
    }

    // CWE-90: LDAP filter built from a request value.
    @GetMapping("/ldap")
    public Object ldap(@RequestParam String uid) throws Exception {
        DirContext context = new InitialDirContext();
        SearchControls controls = new SearchControls();
        controls.setSearchScope(SearchControls.SUBTREE_SCOPE);
        return context.search("dc=northwind,dc=internal", "(uid=" + uid + ")", controls);
    }

    // CWE-601: unvalidated redirect.
    @GetMapping("/redirect")
    public void redirect(@RequestParam String next, HttpServletResponse response) throws Exception {
        response.sendRedirect(next);
    }

    // CWE-644: absolute URL from the Host header.
    @GetMapping("/reset-link")
    public String resetLink(HttpServletRequest request) {
        String host = request.getHeader("Host");
        return "https://" + host + "/auth/reset?token=abc";
    }

    // CWE-200: JWT secret exposed.
    @GetMapping("/debug/jwt")
    public String debugJwt() {
        return jwtService.secret();
    }

    // CWE-352: state-changing route with no CSRF token.
    @PostMapping("/profile")
    public Object profile(@RequestBody Map<String, Object> body) {
        return body;
    }
}
