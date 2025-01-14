import styles from './Header.module.css';

const Header: React.FC = () => {
    return (
        <header className={styles.header}>
            <h1 className={styles.title}>To Do List</h1>
            <nav className={styles.nav}>
                <ul>
                    <li><a href="/" className={styles.link}>Home</a></li>
                    <li><a href="/about" className={styles.link}>About</a></li>
                    <li><a href="/contact" className={styles.link}>Contact</a></li>
                </ul>
            </nav>
        </header>
    );
};

export default Header;