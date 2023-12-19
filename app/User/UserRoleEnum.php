<?php

namespace App\User;

enum UserRoleEnum: string
{
    case ADMIN = 'admin';
    case EDITOR = 'editor';
    case VISITOR = 'visitor';
}
