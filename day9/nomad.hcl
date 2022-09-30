job "uranghishna-echo" {
  datacenters = ["dc1"]
  type = "service"

  group "web" {
    count = 1

    network {
      port "http" {
        to = 1323
      }
    }

    task "uranghishna-echo" {
      driver = "docker"

      config {
        image = "uranghishna/go-echo:v1"
        ports = ["http"]
      }

      resources {
        cpu    = 100
        memory = 128
      }
    }

    service {
      name = "uranghishna-echo"
      port = "http"
      tags = [
        "traefik.enable=true",
        "traefik.http.routers.uranghishna-echo-demo.rule=Host(\"uranghishna.cupang.efishery.ai\")",
      ]
      check {
        port        = "http"
        type        = "tcp"
        interval    = "15s"
        timeout     = "14s"
      }
    }

  }
}