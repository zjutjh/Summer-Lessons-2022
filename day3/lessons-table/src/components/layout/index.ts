import { default as InternalLayout } from "./src/Layout.vue";
import Header from "./src/Header.vue";
import Sider from "./src/Sider.vue";
import Footer from "./src/Footer.vue";
import Content from "./src/Content.vue";

type InternalLayoutType = typeof InternalLayout;

interface LayoutType extends InternalLayoutType {
  Header: typeof Header;
  Footer: typeof Footer;
  Content: typeof Content;
  Sider: typeof Sider;
}
const Layout = InternalLayout as LayoutType;

Layout.Header = Header;
Layout.Footer = Footer;
Layout.Content = Content;
Layout.Sider = Sider;

export default Layout;
