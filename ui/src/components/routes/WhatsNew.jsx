import Markdown from "markdown-to-jsx";
import { MainComponentWrapper } from "../MainComponentWrapper.jsx";
import { useEffect, useState } from "preact/hooks";
import * as whatsNew from "../../assets/announcements/whats-new.md";

export function WhatsNew() {
  const [text, setText] = useState("");

  useEffect(() => {
    try {
      fetch(whatsNew.default)
        .then((res) => res.text())
        .then((importedText) => setText(importedText));
    } catch (e) {
      setText(`Error\n${e.message}`);
    }
  }, []);

  return (
    <MainComponentWrapper useFooter={true}>
      {text !== "" && (
        <article className="prose">
          <Markdown>{text}</Markdown>
        </article>
      )}
    </MainComponentWrapper>
  );
}
