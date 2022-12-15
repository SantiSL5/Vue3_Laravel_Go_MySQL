<?php

// use App\Http\Controllers\BookController;
use App\Http\Controllers\TableController;
use App\Http\Controllers\CategoryController;
use App\Http\Controllers\DishTypeController;
use App\Http\Controllers\DishController;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::middleware('auth:sanctum')->get('/user', function (Request $request) {
    return $request->user();
});

Route::prefix('api/category')->group(function () {
    Route::get('/', [CategoryController::class, 'index']);
    Route::get('/{id}', [CategoryController::class, 'show']);
    Route::post('/', [CategoryController::class, 'store']);
    Route::put('/{id}', [CategoryController::class, 'update']);
    Route::delete('/{id}', [CategoryController::class, 'destroy']);
});

Route::prefix('api/table')->group(function () {
    Route::get('/', [TableController::class, 'index']);
    Route::get('/{id}', [TableController::class, 'show']);
    Route::post('/', [TableController::class, 'store']);
    Route::put('/{id}', [TableController::class, 'update']);
    Route::delete('/{id}', [TableController::class, 'destroy']);
});

Route::prefix('api/dishtype')->group(function () {
    Route::get('/', [DishTypeController::class, 'index']);
    Route::get('/{id}', [DishTypeController::class, 'show']);
    Route::post('/', [DishTypeController::class, 'store']);
    Route::put('/{id}', [DishTypeController::class, 'update']);
    Route::delete('/{id}', [DishTypeController::class, 'destroy']);
});

Route::prefix('api/dish')->group(function () {
    Route::get('/', [DishController::class, 'index']);
    Route::get('/{id}', [DishController::class, 'show']);
    Route::post('/', [DishController::class, 'store']);
    Route::put('/{id}', [DishController::class, 'update']);
    Route::delete('/{id}', [DishController::class, 'destroy']);
});


// Route::get('/mesa', [BookController::class, 'prueba']);
// Route::get('/table', [TableController::class, 'prueba']);
// Route::post('/books', [BookController::class, 'addBook']);
// Route::put('/books/{id}', [BookController::class, 'updateBook']);
// Route::delete('/books/{id}', [BookController::class, 'deleteBook']);
