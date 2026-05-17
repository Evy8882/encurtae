import { useState, useEffect } from "react";
import "./App.css";
import EditUrlModal from "../components/modal";

interface UrlItem {
  id: string;
  originalUrl: string;
  shortUrl: string;
}

function App() {
  const [urlInput, setUrlInput] = useState("");
  const [urls, setUrls] = useState<UrlItem[]>([]);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [editingUrl, setEditingUrl] = useState<UrlItem | null>(null);

  useEffect(() => {
    const storedUrls = localStorage.getItem("urls");
    if (storedUrls) {
      setUrls(JSON.parse(storedUrls));
    }
  }, []);

  async function handleAddUrl() {
    if (!urlInput.trim()) return;

    fetch("http://localhost:8080/post", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ originalUrl: urlInput }),
    })
      .then(async (response) => {
        const text = await response.text();

        if (!response.ok) {
          throw new Error(text);
        }

        return JSON.parse(text);
      })
      .then((data) => {
        const newUrl: UrlItem = {
          id: data.id,
          originalUrl: data.originalUrl,
          shortUrl: data.shortUrl,
        };

        setUrls((prev) => [...prev, newUrl]);
        localStorage.setItem("urls", JSON.stringify([...urls, newUrl]));
        setUrlInput("");
      })
      .catch((error) => {
        console.error("Erro ao encurtar a URL:", error);
      });
  }

  return (
    <>
      <main>
        <h1>Encurtaê</h1>
        <p>Crie links encurtados para suas URL's</p>
        <section>
          <div className="new-url-container">
            <input
              type="text"
              placeholder="Cole sua URL aqui..."
              value={urlInput}
              onChange={(e) => setUrlInput(e.target.value)}
            />
            <button onClick={handleAddUrl}>Adicionar</button>
          </div>
          <h2>Suas URLS:</h2>
          <ul>
            {urls.map((url) => (
              <li key={url.id}>
                <strong>Original:</strong> {url.originalUrl} <br />
                <strong>Encurtada:</strong> {url.shortUrl}
                <div className="actions">
                <button
                  className="dangerous-btn"
                  onClick={() => setUrls(urls.filter((u) => u.id !== url.id))}
                >
                  Excluir
                </button>
                <button className="edit-btn" onClick={() => {
                  setEditingUrl(url);
                  setIsModalOpen(true);
                }}>
                  Editar
                </button>
                </div>
              </li>
            ))}
          </ul>
        </section>
        <EditUrlModal
          isOpen={isModalOpen}
          originalUrl={editingUrl?.originalUrl || ""}
          onClose={() => setIsModalOpen(false)}
          onSave={({ originalUrl }) => {
            if (!editingUrl) return;

            const updatedUrls = urls.map((url) =>
              url.id === editingUrl.id ? { ...url, originalUrl } : url
            );
            setUrls(updatedUrls);
            // localStorage.setItem("urls", JSON.stringify(updatedUrls));
            setIsModalOpen(false);
          }}
        />
      </main>
    </>
  );
}

export default App;
