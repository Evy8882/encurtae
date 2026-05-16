package repository

import (
	"testing"
)

func TestFirebaseConnection(t *testing.T) {
	// Tenta inicializar o serviço que você acabou de criar
	service, err := NewFirebaseService()
	
	// Se der erro, o teste falha e mostra o motivo
	if err != nil {
		t.Fatalf("Falha na conexão com o Firebase: %v", err)
	}

	// Se chegou aqui, o app e o auth foram criados com sucesso
	if service.App == nil {
		t.Error("O objeto Firebase App retornou nulo")
	}
	if service.Auth == nil {
		t.Error("O objeto Firebase Auth retornou nulo")
	}

	t.Log("Conexão com o Firebase estabelecida com sucesso!")
}
