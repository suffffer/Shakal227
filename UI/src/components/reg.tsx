import style from "./reg.module.css";

function RegisterPage() {
    return (
        <div className={style.formContainer}>
            <h2>Register</h2>
            <form action="/register" method="POST">
                <input type="text" name="username" placeholder="Username" required />
                <input type="password" name="password" placeholder="Password" required />
                <input type="password" name="confirm_password" placeholder="Confirm Password" required />
                <div className={style.policy}>
                    <input type="checkbox" name="policy" required />
                    <label htmlFor="policy">I agree to the website policy</label>
                </div>
                <button type="submit">Register</button>
            </form>
        </div>
    );
}

export default RegisterPage;
