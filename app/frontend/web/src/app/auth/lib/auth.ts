"use client";

import { env } from "@/lib/env";
import { saveToken } from "@/app/auth/lib/save-token";

export async function login(email: string, password: string) {
  const res = await fetch(`${env.NEXT_PUBLIC_SERVER_URL}/api/auth/login`, {
    method: "POST",
    body: JSON.stringify({
      email,
      password,
    }),
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
  });

  if (!res.ok) {
    throw new Error(`invalid status code: ${res.status}`);
  }
  saveToken(res.headers.get("Authorization") ?? "");
  window.location.href = "/home";
}

export async function register(email: string, password: string) {
  const res = await fetch(`${env.NEXT_PUBLIC_SERVER_URL}/api/auth/register`, {
    method: "POST",
    body: JSON.stringify({
      user: {
        email,
        password,
      },
    }),
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
  });

  if (!res.ok) {
    throw new Error(`invalid status code: ${res.status}`);
  }
  saveToken(res.headers.get("Authorization") ?? "");
  window.location.href = "/home";
}
