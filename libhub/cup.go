package libhub

import (
  "regexp"
)

type repo struct {
  User string
  RepoName string
  Branch string
}

type repoInfo struct {
  Default_branch string `json: "default_branch"`
  Parent struct {
    Name string `json: "parent.name"`
    Owner struct {
        Login string `json: "parent.owner.login"`
      } `json: "parent.owner"`
  } `json: "parent"`
}

func parseRepo(str string, omitUser bool) (r repo) {
  rgx := regexp.MustCompile(`/`)
  strs := rgx.Split(str, 3)
  switch len(strs) {
  case 0:
    return
  case 1:
    if omitUser {
      r.RepoName = strs[0]
    } else {
      r.User = strs[0]
    }
  case 2:
    if omitUser {
      r.RepoName = strs[0]
      r.Branch = strs[1]
    } else {
      r.User = strs[0]
      r.RepoName = strs[1]
    }
  case 3:
    r.User = strs[0]
    r.RepoName = strs[1]
    r.Branch = strs[2]
  }
  return
}

func (hc *hubCup) Cup(what string, from string, force bool, dryRun bool) error {
  var err error
  logger.Debugf("Parsing repos...")
  whatRepo := parseRepo(what, true)
  fromRepo := parseRepo(from, false)
  if len(whatRepo.User) == 0 {
    whatRepo.User, err = hc.getMe()
    if err != nil {
      return err
    }
  }
  logger.Debugf("Parsed what:%+v; Parsed from:%+v;", whatRepo, fromRepo)
  if len(whatRepo.Branch) == 0 || len(fromRepo.User) == 0 {
    wrif, err := hc.getRepo(whatRepo)
    if err != nil {
      return err
    }
    logger.Debugf("Get repo:%+v", wrif)
    if len(whatRepo.Branch) == 0 {
      whatRepo.Branch = wrif.Default_branch
    }
    if len(fromRepo.User) == 0 {
      fromRepo.User = wrif.Parent.Owner.Login
      fromRepo.RepoName = wrif.Parent.Name
    }
  }
  if len(fromRepo.RepoName) == 0 {
    fromRepo.RepoName = whatRepo.RepoName
  }
  if len(fromRepo.Branch) == 0 {
    fromRepo.Branch = whatRepo.Branch
  }
  logger.Debugf("Filled what:%+v; Filled from:%+v;", whatRepo, fromRepo)
  
  if dryRun {
    logger.Infof("Stopped because of --dry-run.")
    return nil
  }

  return nil
}
