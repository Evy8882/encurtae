import './App.css'

function App() {

  return (
    <>
      <main>
        <h1>Encurtaê</h1>
        <p>Crie links encurtados para suas URL's</p>
        <section>
          <div className='new-url-container'>
            <input type="text" placeholder="Cole sua URL aqui..." />
            <button>Adicionar</button>
          </div>
          <h2>Suas URLS:</h2>
          <ul></ul>
        </section>
      </main>
    </>
  )
}

export default App
