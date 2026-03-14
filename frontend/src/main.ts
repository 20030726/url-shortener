import "./style.css";
import { shortenURL } from "./api.ts";

interface HistoryEntry {
  original: string;
  short: string;
}

const form = document.getElementById("shorten-form") as HTMLFormElement;
const input = document.getElementById("url-input") as HTMLInputElement;
const errorMsg = document.getElementById("error-msg") as HTMLParagraphElement;
const resultDiv = document.getElementById("result") as HTMLDivElement;
const shortUrlAnchor = document.getElementById("short-url") as HTMLAnchorElement;
const copyBtn = document.getElementById("copy-btn") as HTMLButtonElement;
const historyList = document.getElementById("history") as HTMLUListElement;
const shortenBtn = document.getElementById("shorten-btn") as HTMLButtonElement;

const history: HistoryEntry[] = [];

form.addEventListener("submit", async (e: Event) => {
  e.preventDefault();
  const longURL = input.value.trim();
  if (!longURL) return;

  errorMsg.textContent = "";
  resultDiv.classList.add("hidden");
  shortenBtn.disabled = true;
  shortenBtn.textContent = "Shortening…";

  try {
    const data = await shortenURL(longURL);
    shortUrlAnchor.textContent = data.short_url;
    shortUrlAnchor.href = data.short_url;
    resultDiv.classList.remove("hidden");

    history.unshift({ original: longURL, short: data.short_url });
    renderHistory();

    input.value = "";
  } catch (err: unknown) {
    errorMsg.textContent =
      err instanceof Error ? err.message : "An unexpected error occurred.";
  } finally {
    shortenBtn.disabled = false;
    shortenBtn.textContent = "Shorten";
  }
});

copyBtn.addEventListener("click", async () => {
  const url = shortUrlAnchor.textContent ?? "";
  if (!url) return;
  try {
    await navigator.clipboard.writeText(url);
    copyBtn.textContent = "Copied!";
    copyBtn.classList.add("copied");
    setTimeout(() => {
      copyBtn.textContent = "Copy";
      copyBtn.classList.remove("copied");
    }, 2000);
  } catch {
    copyBtn.textContent = "Failed";
  }
});

function renderHistory(): void {
  historyList.innerHTML = "";
  if (history.length === 0) return;

  const heading = document.createElement("li");
  heading.className = "history-heading";
  heading.textContent = "Recent links";
  historyList.appendChild(heading);

  history.slice(0, 10).forEach((entry) => {
    const li = document.createElement("li");
    li.className = "history-item";

    const shortLink = document.createElement("a");
    shortLink.href = entry.short;
    shortLink.target = "_blank";
    shortLink.rel = "noopener noreferrer";
    shortLink.textContent = entry.short;
    shortLink.className = "history-short";

    const original = document.createElement("span");
    original.textContent = entry.original;
    original.className = "history-original";

    li.appendChild(shortLink);
    li.appendChild(original);
    historyList.appendChild(li);
  });
}
