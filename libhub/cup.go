package libhub

import (
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

func parseRepo(str string) (r repo, err error) {
  // TODO
  return repo{User:"", RepoName:"GoAuthing", Branch:""}, nil
}

func (hc *hubCup) Cup(what string, from string) error {
  logger.Debugf("Parsing repos...")
  whatRepo, err := parseRepo(what)
  if err != nil {
    return err
  }
  fromRepo, err := parseRepo(from)
  if err != nil {
    return err
  }
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
  return nil
}
