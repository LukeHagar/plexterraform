terraform {
  required_providers {
    PlexAPI = {
      source  = "LukeHagar/PlexAPI"
      version = "0.15.0"
    }
  }
}

provider "PlexAPI" {
  # Configuration options
}