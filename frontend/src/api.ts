export interface ShortenResponse {
  short_url: string;
}

export async function shortenURL(longURL: string): Promise<ShortenResponse> {
  const response = await fetch("/shorten", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ url: longURL }),
  });

  if (!response.ok) {
    const text = await response.text();
    throw new Error(text.trim() || `Request failed (${response.status})`);
  }

  return response.json() as Promise<ShortenResponse>;
}
