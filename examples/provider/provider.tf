terraform {
  required_providers {
    PlexAPI = {
      source  = "LukeHagar/PlexAPI"
      version = "0.2.1"
    }
  }
}

provider "PlexAPI" {
  # Configuration options
}