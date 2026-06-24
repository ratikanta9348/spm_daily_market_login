document.addEventListener("DOMContentLoaded", () => {

    const loginForm = document.getElementById("loginForm");
    const errorDiv = document.getElementById("error");

    loginForm.addEventListener("submit", async (e) => {

        e.preventDefault();

        errorDiv.innerText = "";

        const userID = document.getElementById("userid").value.trim();
        const password = document.getElementById("password").value.trim();

        if (!userID || !password) {
            errorDiv.innerText = "Please enter User ID and Password";
            return;
        }

        const formData = new URLSearchParams();
        formData.append("userid", userID);
        formData.append("password", password);

        try {

            const response = await fetch("/authenticate", {
                method: "POST",
                headers: {
                    "Content-Type": "application/x-www-form-urlencoded"
                },
                body: formData
            });

            const data = await response.json();

            if (data.success) {
                
                errorDiv.style.color = "#22c55e";
                errorDiv.innerText = "Login successful. Redirecting...";

                setTimeout(() => {
                                   errorDiv.remove();
                    window.location.href = "https://creatorapp.zohopublic.in/maamangalaenterprises/spm/form-perma/Daily_Fish_Market_Price_Entry/sBY2pU8zj4UMO6t5nhde53ghXYQYAQFHWhWt3ptbVRF3YTuKE3sQNN3NyaDraUn2fGQbpP519ZQd4yGqUwTayTOK7kMVSxRWKdfP";
                }, 1000);

            } else {

                errorDiv.style.color = "#ef4444";
                errorDiv.innerText =
                    data.message || "Invalid User ID or Password";
            }

        } catch (err) {

            console.error(err);

            errorDiv.style.color = "#ef4444";
            errorDiv.innerText =
                "Unable to connect to server";
        }
    });
});