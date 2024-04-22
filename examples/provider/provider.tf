terraform {
  required_providers {
    PlexAPI = {
      source  = "LukeHagar/PlexAPI"
      version = "0.11.2"
    }
  }
}

provider "PlexAPI" {
  # Configuration options
}