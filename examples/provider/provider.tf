terraform {
  required_providers {
    PlexAPI = {
      source  = "LukeHagar/PlexAPI"
      version = "0.16.0"
    }
  }
}

provider "PlexAPI" {
  # Configuration options
}