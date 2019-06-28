workflow "Mirror to GitLab" {
  on = "push"
  resolves = ["Push to GitLab"]
}

action "Push to GitLab" {
  uses = "spyoungtech/mirror-action@master"
  args = "https://gitlab.com/dryproject/drycop.git"
  env = {
    GIT_USERNAME = "artob"
  }
  secrets = ["GIT_PASSWORD"]
}
