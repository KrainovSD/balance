export function copyToClipboard(textToCopy: string) {
  return new Promise<boolean>((resolve) => {
    try {
      if (navigator.clipboard && window.isSecureContext) {
        navigator.clipboard
          .writeText(textToCopy)
          .then(() => resolve(true))
          .catch(() => resolve(false));

        return;
      }

      const textArea = document.createElement("textarea");
      textArea.value = textToCopy;
      textArea.style.position = "fixed";
      textArea.style.left = "-999999px";
      textArea.style.top = "-999999px";
      document.body.appendChild(textArea);
      textArea.focus();
      textArea.select();

      // eslint-disable-next-line @typescript-eslint/no-deprecated
      const result = document.execCommand("copy");
      textArea.remove();

      resolve(result);
    } catch {
      resolve(false);
    }
  });
}
