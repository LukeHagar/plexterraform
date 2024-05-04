terraform {
  required_providers {
    PlexAPI = {
      source  = "LukeHagar/PlexAPI"
      version = "0.12.0"
    }
  }
}

provider "PlexAPI" {
  # Configuration options
}