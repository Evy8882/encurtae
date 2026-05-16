import { useState } from 'react'
import './App.css'

interface UrlItem {
  id: string
  originalUrl: string
  shortUrl: string
}

function App() {
  const [urlInput, setUrlInput] = useState('')
  const [urls, setUrls] = useState<UrlItem[]>([{
  id: '1',
  originalUrl: 'https://www.exemplo.com/uma-url-muito-longa',
  shortUrl: 'http://encurtae.com/abc123'
  }])

  function handleAddUrl() {
    if (!urlInput.trim()) return
    // Adicionar a requisição para o backend
    const newUrl: UrlItem = {
      id: Date.now().toString(),
      originalUrl: urlInput,
      shortUrl: `http://encurtae.com/${Math.random().toString(36).substring(7)}`
    }
    setUrls([...urls, newUrl])
    setUrlInput('')
  }

  return (
    <>
      <main>
        <h1>Encurtaê</h1>
        <p>Crie links encurtados para suas URL's</p>
        <section>
          <div className='new-url-container'>
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
                <button className='dangerous-btn' onClick={() => setUrls(urls.filter((u) => u.id !== url.id))}>
                  Excluir
                </button>
              </li>
            ))}
          </ul>
        </section>
      </main>
    </>
  )
}

export default App
