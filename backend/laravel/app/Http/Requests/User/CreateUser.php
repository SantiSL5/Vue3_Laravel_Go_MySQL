<?php

namespace App\Http\Requests\User;

use Illuminate\Foundation\Http\FormRequest;
use Illuminate\Http\Exceptions\HttpResponseException;

class CreateUser extends FormRequest {
    
    public function authorize() {
        return true;
    }

    public function rules() {
        return [
            'username' => 'unique:users|required',
            'email' => 'unique:users|required',
            'password' => 'required',
            'type' => 'sometimes',
            'image' => 'sometimes',
            'enabled' => 'sometimes',
        ];
    }

    public function withValidator($validator) {
        if ($validator->fails()) {
            throw new HttpResponseException(response()->json([
                'msg' => 'Introduce correctamente los campos para crear un usuario',
                'status' => false,
                'errors' => $validator->errors(),
            ], 400));
       }
    }
}