<?php

namespace App\Enum;

enum UserRoleEnum: string
{
    case ADMIN = 'admin';
    case EDITOR = 'editor';
    case VISITOR = 'visitor';
}
