<?php

namespace App\Services;

use App\Interfaces\UserRepositoryInterface;
use App\Http\Requests\User\CreateUser;
use App\Http\Requests\User\UpdateUser;
use App\Http\Requests\User\LoginUser;
use App\Http\Resources\UserResource;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\Hash;

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
            if ($request->get('password')) {
                $passwordhashed=bcrypt($request->get('password'));
                $request->merge([
                    'password' => $passwordhashed,
                ]);
            }
            $update = $this->userRepository->updateUserRepo($request, $id);
            if ($update == 1) {
                return UserResource::make($this->userRepository->getUserByIdRepo($id));
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
        $credentials = $request->only('email', 'password');
        $token=Auth::attempt($credentials);
        if (!$token) {
            return response()->json([ "error" => "Unauthorized" ], 400);
        }

        try {
            $user=UserResource::make($this->userRepository->getUserByEmailRepo($request->get('email')));
            if ($user['type'] == "admin") {
                return response()->json(['token' => $token, 'user' => $this->userRepository->loginUserRepo()]);
            }
            return response()->json([
                "Message" => "This user is not an admin"
            ]);
        } catch (\Exception $e) {
            return response()->json([
                "Message" => "User doesn't exist"
            ]);
        }
    }

    public function profileUserService()
    {
        $user_serialize = Auth::user();
        if ($user_serialize == null) {
            return response()->json([
                "Message" => "Invalid token"
            ]);
        }
        if ($user_serialize->type != "admin") {
            return response()->json([
                "Message" => "This user is not an admin"
            ]);
        }
        unset($user_serialize->password);
        return response()->json($user_serialize);
    }
}