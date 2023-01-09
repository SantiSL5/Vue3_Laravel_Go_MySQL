<?php

namespace App\Interfaces;
use App\Http\Requests\User\CreateUser;
use App\Http\Requests\User\UpdateUser;
use App\Http\Requests\User\LoginUser;

interface UserRepositoryInterface
{
    function getAllUsersRepo();
    function getUserByIdRepo($id);
    function getUserByUsernameRepo($username);
    function createUserRepo(CreateUser $request);
    function loginUserRepo(LoginUser $request);
    function updateUserRepo(UpdateUser $request,$id);
    function deleteUserRepo($id);
}
