# -*- coding: utf-8 -*-

from invoke import task


@task
def clean(ctx):
    ctx.run("find . -path ./.git -prune -o -type f -perm 755 -print -exec rm -rf {} +", echo=True)
