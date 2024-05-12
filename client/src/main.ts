import "bootstrap";
import "bootstrap/scss/bootstrap.scss";

document.getElementById("generate")?.addEventListener("click", async () => {
  const res = await fetch("/api/v1/passwords");
  const pass = await res.text();

  document.getElementById("password")!.textContent = pass;
});
