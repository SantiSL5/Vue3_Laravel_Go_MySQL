<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Http\Requests\StoreUserRequest;
use App\Http\Requests\User\CreateUser;
use App\Http\Requests\User\UpdateUser;
use App\Http\Requests\User\LoginUser;
use App\Services\UserService;
use App\Http\Resources\UserResource;

use Illuminate\Support\Facades\Auth;

class UserController extends Controller {


    public function __construct(UserService $userService) {
        $this->userService = $userService;
    }

    public function index() 
    {
        return $this->userService->getAllUsersService();
    }

    /**
     * Show the form for creating a new resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function create()
    {
        //
    }

    /**
     * Store a newly created resource in storage.
     *
     * @param  \App\Http\Requests\User\CreateUser  $request
     * @return \Illuminate\Http\Response
     */
    public function store(CreateUser $request) 
    {
        return $this->userService->createUserService($request);
    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  \App\Models\User  $table
     * @return \Illuminate\Http\Response
     */
    public function edit(User $user)
    {
        //
    }

    /**
     * Display the specified resource.
     *
     * @param  \App\Models\User $user
     * @return \Illuminate\Http\Response
     */
    public function show($id) 
    {
        return $this->userService->getUserByIdService($id);
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \App\Http\Requests\User\UpdateUser  $request
     * @param  \App\Models\User  $user
     * @return \Illuminate\Http\Response
     */
    public function update(UpdateUser $request, $id) 
    {
        return $this->userService->updateuserService($request,$id);
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  \App\Models\User  $user
     * @return \Illuminate\Http\Response
     */

    public function destroy($id) 
    {
        return $this->userService->deleteUserService($id);
    }

    public function login(LoginUser $request) 
    {
        return $this->userService->loginUserService($request);
    }

}