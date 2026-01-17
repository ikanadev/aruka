import { useEffect } from "react";
import styles from "./ChatTextAreaInput.module.css";

const MAX_TEXT_AREA_HEIGHT = 250;

interface Props {
  textAreaRef: React.RefObject<HTMLTextAreaElement | null>;
}
export function ChatTextAreaInput(props: Props) {
  const { textAreaRef } = props;

  useEffect(() => {
    const textarea = textAreaRef.current;
    if (!textarea) return;
    adjustHeight(textarea);
    const adjustHeightCb = () => adjustHeight(textarea);
    textarea.addEventListener("input", adjustHeightCb);
    return () => {
      textarea.removeEventListener("input", adjustHeightCb);
    };
  }, [textAreaRef]);

  return <textarea ref={textAreaRef} className={styles.textArea} placeholder="Ask anything..." />;
}

function adjustHeight(textarea: HTMLTextAreaElement) {
  textarea.style.height = "auto";
  const newHeight = Math.min(textarea.scrollHeight, MAX_TEXT_AREA_HEIGHT);
  textarea.style.height = `${newHeight}px`;
  textarea.style.overflowY = textarea.scrollHeight > MAX_TEXT_AREA_HEIGHT ? "auto" : "hidden";
}
