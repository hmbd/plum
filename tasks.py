# -*- coding: utf-8 -*-

from invoke import task


@task
def clean(ctx):
    ctx.run("find . -path ./.git -prune -o -perm 755 -type f | grep -v "./.git" -exec rm -rf {} +", echo=True)
