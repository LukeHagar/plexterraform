terraform {
  required_providers {
    PlexAPI = {
      source  = "LukeHagar/PlexAPI"
      version = "0.2.5"
    }
  }
}

provider "PlexAPI" {
  # Configuration options
}