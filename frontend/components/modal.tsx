import React from "react";

interface EditUrlModalProps {
  isOpen: boolean;
  originalUrl: string;
  onClose: () => void;
  onSave: (payload: { originalUrl: string }) => void;
}

function EditUrlModal({
  isOpen,
  originalUrl,
  onClose,
  onSave,
}: EditUrlModalProps) {
  const [originalValue, setOriginalValue] = React.useState(originalUrl);

  React.useEffect(() => {
    if (isOpen) {
      setOriginalValue(originalUrl);
    }
  }, [isOpen, originalUrl]);

  if (!isOpen) return null;

  return (
    <div className="modal-overlay" role="dialog" aria-modal="true">
      <div className="modal-card">
        <header className="modal-header">
          <h3>Editar link</h3>
          <button className="modal-close" onClick={onClose} aria-label="Fechar">
            ×
          </button>
        </header>

        <div className="modal-body">
          <label htmlFor="originalUrl">URL original</label>
          <input
            type="text"
            id="originalUrl"
            value={originalValue}
            onChange={(e) => setOriginalValue(e.target.value)}
          />
        </div>

        <footer className="modal-actions">
          <button className="modal-secondary" onClick={onClose}>
            Cancelar
          </button>
          <button
            className="modal-primary"
            onClick={() => onSave({ originalUrl: originalValue })}
          >
            Salvar
          </button>
        </footer>
      </div>
    </div>
  );
}

export default EditUrlModal;
