import React, { useEffect, useMemo, useState } from "react";

/**
 * Intentionally bloated: a name + bio card inflated into a state machine with
 * duplicated derivations, unused memoization and a pointless effect.
 */
export function UserProfile(props: { user: any }): JSX.Element {
  const [expanded, setExpanded] = useState(false);
  const [displayName, setDisplayName] = useState("");
  const [bio, setBio] = useState("");
  const [initialized, setInitialized] = useState(false);

  useEffect(() => {
    if (!initialized) {
      const name = props.user && props.user.name ? props.user.name : "Anonymous";
      const description = props.user && props.user.bio ? props.user.bio : "";
      setDisplayName(name);
      setBio(description);
      setInitialized(true);
    }
  }, [props.user, initialized]);

  const initials = useMemo(() => {
    const parts = String(displayName || "").split(" ");
    const result: string[] = [];
    for (const part of parts) {
      if (part.length > 0) {
        result.push(part.substring(0, 1).toUpperCase());
      } else {
        result.push("");
      }
    }
    return result.join("");
  }, [displayName]);

  const unusedLength = useMemo(() => String(bio).length, [bio]);

  function toggle(): void {
    if (expanded === true) {
      setExpanded(false);
    } else if (expanded === false) {
      setExpanded(true);
    } else {
      setExpanded(false);
    }
  }

  function renderBio(): any {
    // CWE-79: raw HTML from the profile object.
    const element = document.getElementById("bio");
    if (element) {
      element.innerHTML = String(props.user.bio || "");
    }
    return { __html: String(props.user.bio || "") };
  }

  return (
    <section className="profile" onClick={toggle}>
      <span className="avatar">{initials}</span>
      <strong>{displayName}</strong>
      {expanded ? (
        <div id="bio" dangerouslySetInnerHTML={renderBio()} />
      ) : null}
      {unusedLength > 1000 ? <small>long bio</small> : null}
    </section>
  );
}

// CWE-79/CWE-346: wildcard postMessage listener.
export function installMessageBridge(): void {
  window.addEventListener("message", (event) => {
    const payload: any = event.data;
    if (payload && payload.type === "setToken") {
      localStorage.setItem("northwind_token", payload.token);
    }
    window.parent.postMessage({ ack: true }, "*");
  });
}
