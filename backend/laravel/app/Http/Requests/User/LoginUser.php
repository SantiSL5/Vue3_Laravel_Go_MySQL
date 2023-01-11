<?php

namespace App\Http\Requests\User;

use Illuminate\Foundation\Http\FormRequest;
use Illuminate\Http\Exceptions\HttpResponseException;

class LoginUser extends FormRequest {
    
    public function authorize() {
        return true;
    }

    public function rules() {
        return [
            'email' => 'required',
            'password' => 'required',
        ];
    }

    public function withValidator($validator) {
        if ($validator->fails()) {
            throw new HttpResponseException(response()->json([
                'msg' => 'Introduce correctamente los campos para loguearte',
                'status' => false,
                'errors' => $validator->errors(),
            ], 400));
       }
    }
}