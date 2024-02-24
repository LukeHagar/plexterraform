terraform {
  required_providers {
    PlexAPI = {
      source  = "LukeHagar/PlexAPI"
      version = "0.6.1"
    }
  }
}

provider "PlexAPI" {
  # Configuration options
}