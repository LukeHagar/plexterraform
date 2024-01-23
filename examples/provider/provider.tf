terraform {
  required_providers {
    PlexAPI = {
      source  = "LukeHagar/PlexAPI"
      version = "0.3.3"
    }
  }
}

provider "PlexAPI" {
  # Configuration options
}