FROM williamyeh/ansible:ubuntu14.04-onbuild

# ==> Specify playbook filename;   default = "playbook.yml"
ENV PLAYBOOK   home.yml

# ==> Specify inventory filename;  default = "/etc/ansible/hosts"
#ENV INVENTORY  inventory.ini

# ==> Executing Ansible...
RUN ansible-playbook-wrapper
