<?php

namespace App\Services;

use App\Interfaces\UserRepositoryInterface;
use App\Http\Requests\User\CreateUser;
use App\Http\Requests\User\UpdateUser;
use App\Http\Requests\User\LoginUser;
use App\Http\Resources\UserResource;

class UserService 
{
    private $userRepository;

    public function __construct(UserRepositoryInterface $userRepository)
    {
        $this->userRepository = $userRepository;
    }

    public function getAllUsersService()
    {
        return UserResource::collection($this->userRepository->getAllUsersRepo());
    }

    public function getUserByIdService($id) 
    {
        try {
            return UserResource::make($this->userRepository->getUserByIdRepo($id));
        } catch (\Exception $e) {
            return response()->json([
                "Message" => "User doesn't exist"
            ]);
        }
    }

    public function createUserService(CreateUser $request)
    {
        try {
            $passwordhashed=bcrypt($request->get('password'));
            $request->merge([
                'password' => $passwordhashed,
            ]);
            // return $request;
            return UserResource::make($this->userRepository->createUserRepo($request));
        } catch (\Exception $e) {
            return response()->json([
                "Message" => "User already exist"
            ]);
        }
    }

    public function updateUserService(UpdateUser $request, $id)
    {
        try {
            UserResource::make($this->userRepository->getUserByIdRepo($id));
        } catch (\Exception $e) {
            return response()->json([
                "Message" => "User doesn't exist"
            ]);
        }

        try {
            $update = $this->userRepository->updateUserRepo($request, $id);
            if ($update == 1) {
                return response()->json([
                    "Message" => "Updated correctly"
                ]);
            } else {
                return response()->json([
                    "Message" => "Not found"
                ], 404);
            };
        } catch (\Exception $e) {
            return response()->json("User already exists");
        }
    }

    public function deleteUserService($id)
    {
        $delete = $this->userRepository->deleteUserRepo($id);

        if ($delete == 1) {
            return response()->json([
                "Message" => "Deleted correctly"
            ], 200);
        } else {
            return response()->json([
                "Status" => "Not found"
            ], 404);
        } 
    }

    public function loginUserService(LoginUser $request)
    {
        $token = Auth::attempt($request->validated());

        if (!$token) {
            return response()->json([ "error" => "Unauthorized" ], 400);
        }

        try {
            return UserResource::make($this->userRepository->getUserByUsernameRepo($request->get('username')));
            return response()->json(['token' => $token, 'user' => $this->userRepository->loginUserRepo()]);
        } catch (\Exception $e) {
            return response()->json([
                "Message" => "User doesn't exist"
            ]);
        }
    }
}