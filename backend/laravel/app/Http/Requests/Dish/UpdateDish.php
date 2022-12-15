<?php

namespace App\Http\Requests\Dish;

use Illuminate\Foundation\Http\FormRequest;

class UpdateDish extends FormRequest
{
    /**
     * Determine if the user is authorized to make this request.
     *
     * @return bool
     */
    public function authorize()
    {
        return true;
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array
     */
    public function rules()
    {
        return [
            'dish' => ['sometimes'],
            'type' => ['sometimes'],
            'price' => ['sometimes'],
            'desc' => ['sometimes'],
            'photo' => ['sometimes'],
        ];
    }
}
