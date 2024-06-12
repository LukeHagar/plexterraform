terraform {
  required_providers {
    PlexAPI = {
      source  = "LukeHagar/PlexAPI"
      version = "0.15.3"
    }
  }
}

provider "PlexAPI" {
  # Configuration options
}