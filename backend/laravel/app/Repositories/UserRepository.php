<?php

namespace App\Repositories;
use App\Models\User;
use App\Interfaces\UserRepositoryInterface;
use App\Http\Requests\User\CreateUser;
use App\Http\Requests\User\UpdateUser;
use App\Http\Requests\User\LoginUser;
use Illuminate\Database\Eloquent\Builder;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Support\Facades\DB;

class UserRepository implements UserRepositoryInterface
{

    public function getAllUsersRepo()
    {
        return User::all();
    }

    public function getUserByIdRepo($id) 
    {
        return User::where('id', $id)->firstOrFail();
    }

    public function getUserByUsernameRepo($username) 
    {
        return User::where('username', $username)->firstOrFail();
    }

    public function createUserRepo(CreateUser $request)
    {
        return User::create($request->all());
    }

    public function loginUserRepo(LoginUser $request)
    {
        return User::make(auth()->user());
    }

    public function updateUserRepo(UpdateUser $request, $id)
    {
        return User::where('id', $id)->update($request->validated());
    }

    public function deleteUserRepo($id)
    {
        return User::where('id', $id)->delete();
    }
}
