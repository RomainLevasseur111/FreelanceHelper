import React, { useState } from "react";
import { useForm } from "react-hook-form";

function RegisterForm() {
    const {
        register,
        handleSubmit,
        formState: { errors },
        setError,
    } = useForm();
    const onSubmit = (data) =>
        fetch("api/register", {
            method: "POST",
            body: JSON.stringify(data),
        })
            .then(async (response) => {
                if (!response.ok) {
                    return setError("root", {
                        type: "server",
                        message: await response.json(),
                    });
                } else {
                    console.log(await response.json());
                }
            })
            .then()
            .catch(() =>
                setError("root", {
                    type: "server",
                    message:
                        "Sorry, we cannot register your profile now. Try again later.",
                })
            );

    return (
        <form onSubmit={handleSubmit(onSubmit)}>
            <span>{errors.root?.message}</span>
            <input
                type="text"
                placeholder="Name"
                {...register("Name", { required: false, maxLength: 80 })}
            />
            <input
                type="text"
                placeholder="Email"
                {...register("Email", {
                    required: true,
                    pattern: /^\S+@\S+$/i,
                })}
            />
            <input
                type="password"
                placeholder="Password"
                {...register("Password", { required: true })}
            />

            <input type="submit" />
        </form>
    );
}

export default RegisterForm;
