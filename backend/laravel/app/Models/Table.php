<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use App\Http\Models\Category;

class Table extends Model
{
    use HasFactory;
    protected $fillable = [
        'code',
        'category',
        'capacity',
        'reserved'
    ];

    public function tables()
    {
        return $this->belongsTo(Category::class);
    }

}   
