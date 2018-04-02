# ReadMe

Collection of useful commands related to Git

**Initialize Git repo for the first time**

```git
$   git init
```

**Add all content for tracking in the branch**

```git
$   git add .
```

**Commit with message**

```git
$   git commit -m "Initial Commit"
```

**Create new branch**

```git
$   git branch phase1
```

**Change branch to work in differet one**

```git
$   git checkout phase1
```

**Add Remote branch**

```git
$ git remote add learningGo https://github.com/bootstrapsp/learningGo.git
```

**Push all to the remote branch**

```git
git push --set-upstream learningGo phase1
```

**Pull all updates from remote branch**

```git
git fetch --all
```

**Discard all local changes and merge all updates fetched from remote branch**

```git
git reset --hard origin/development
```

**Pulling all the changes from the the remote repo**

```git
git pull
```

**Following used for syncing the github's repo with Bitbucket by adding it to the remote list**

```git
git remote set-url learningGo  --add https://bootstrapsp@bitbucket.org/rajtheceo/learninggo.git
```

**To merge a branch with master**
```git
git merge phase3
```
Note: this should be executed from master branch, i.e. prior to this do git checkout master to change the branch

**Then push to master for final merge**
```git
git push --set-upstream --force learningGo master
```
Note: here learningGo is remote repo name