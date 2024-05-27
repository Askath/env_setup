import os
import subprocess

def update_git_repos(repos_dir='./repos'):
    if not os.path.exists(repos_dir):
        print(f"Directory {repos_dir} does not exist.")
        return

    for root, dirs, files in os.walk(repos_dir):
        if '.git' in dirs:
            repo_path = os.path.abspath(root)
            print(f"Updating repository: {repo_path}")
            try:
                subprocess.run(['git', '-C', repo_path, 'pull'], check=True)
            except subprocess.CalledProcessError as e:
                print(f"Failed to update repository {repo_path}: {e}")
            # To avoid recursion into nested git repos
            dirs[:] = []

if __name__ == '__main__':
    update_git_repos()
