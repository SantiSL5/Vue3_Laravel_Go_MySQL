<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Category extends Model
{
    use HasFactory;
    protected $fillable = [
        'name'
    ];

    // public function mesas(): BelongsToMany
    // {
    //     return $this->belongsToMany(Mesa::class, 'mesas_categories', 'category_id', 'mesa_id');
    // }

}
