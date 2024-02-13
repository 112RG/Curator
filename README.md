<h1 align="center">Curator</h1>
<div align="center"

[![license](https://img.shields.io/github/license/112RG/Curator)](https://github.com/112RG/Curator/blob/master/LICENSE)
![GitHub repo size](https://img.shields.io/github/repo-size/112RG/Curator)
![Lines of Code](https://aschey.tech/tokei/github/112RG/Curator)
</div>

# Curator - FlyPastebin

Curator is a lightweight and fast pastebin application deployed on [Fly.io](https://fly.io/). It allows users to quickly share and store text snippets with ease. This project is built on a simple and efficient stack, utilizing [LiteFS](https://github.com/superfly/litefs) as its database and written in Go.

## Features

- **Fast and Lightweight:** Built for speed and efficiency, ensuring a seamless experience for users.
- **Easy Deployment:** Deploy your own instance on Fly.io effortlessly.
- **Syntax Highlighting:** Supports syntax highlighting for various programming languages.
- **Expiration:** Set expiration time for pastes, keeping your pastebin clean and clutter-free.
- **LiteFS Database:** Utilizes Sqlite LiteFS for efficient and lightweight storage.

## Getting Started

### Prerequisites

1. [Fly.io Account](https://fly.io/) - Sign up for a Fly.io account if you don't have one.
2. [Flyctl CLI](https://fly.io/docs/flyctl/install/) - Install the Flyctl command-line tool.
3. [Go](https://golang.org/doc/install) - Install Go to build and run the Go application.

### Deployment

Clone this repository:

   ```bash
   git clone https://github.com/112RG/Curator.git
   cd Curator
   ```

Build and run Curator for fly.io:

Open a new terminal and navigate to the project directory. Initialize and deploy using Flyctl:
```bash
flyctl init
flyctl deploy
```
Follow the prompts to deploy your Curator instance on Fly.io.
Once deployed, open the pastebin in your browser:
```bash
flyctl open
```


### Development

Clone this repository:

   ```bash
   git clone https://github.com/112RG/Curator.git
   cd Curator
   ```

Build and run Curator for development:

```bash
go run .\main.go -dsn curator.db --addr :8111
```

## Database (Sqlite)(LiteFS)

Curator uses Sqlite using [LiteFS](https://github.com/superfly/litefs) as its database.

## Contributing

Feel free to contribute to the project by opening issues or submitting pull requests. We welcome any improvements, bug fixes, or new features!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by the simplicity of pastebin services.
- Built with Go, [LiteFS](https://github.com/superfly/litefs), and [Fly.io](https://fly.io/).

---

Happy pasting with Curator! 🚀
