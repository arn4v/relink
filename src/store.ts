import { writable } from "svelte/store";
import { API_BASE_URL } from "./constants";

export const data = writable({});

data.subscribe((value) => {
  fetch(API_BASE_URL + "/data", {
    method: "POST",
    body: JSON.stringify(value),
  })
    .then((res) => res.json())
    .then((apiData) => data.set(apiData));
});
