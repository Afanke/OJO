import {
  Pagination,
  Dialog, 
  Menu,
  Submenu,
  MenuItem,
  MenuItemGroup,
  Input,
  InputNumber,
  Container,
  Header,
  Aside,
  Main,
  Footer,
  Link,
  Divider,
  Image,
  Loading,
  Message,
  Drawer,
  Switch,
  Select,
  Option,
  OptionGroup,
  Button,
  ButtonGroup,
  Table,
  TableColumn,
  Icon,
  Row,
  Col,
  // Card,
  Progress,
  Form,
  FormItem,
  // Tabs,
  // TabPane,
  // Tag,
  // Tree,
  // Alert,
  // Slider,

  // Upload,

  // Spinner,
  // Badge,

  // Rate,
  // Steps,
  // Step,
  // Carousel,
  // CarouselItem,
  // Collapse,
  // CollapseItem,
  // Cascader,
  // ColorPicker,
  // Transfer,
  DatePicker,
  // TimeSelect,
  // TimePicker,
  Popover,
  Tooltip,
  // Breadcrumb,
  // BreadcrumbItem,
  Radio,
  // RadioGroup,
  // RadioButton,
  Checkbox,
  CheckboxButton,
  CheckboxGroup,
  // Autocomplete,
  // Dropdown,
  // DropdownMenu,
  // DropdownItem,
  // Timeline,
  // TimelineItem,

 
  // Calendar,
  // Backtop,
  PageHeader,
  // CascaderPanel,
  // MessageBox,
  // Notification,
  Avatar,
  Tag
} from 'element-ui';
import lang from 'element-ui/lib/locale/lang/en'
import locale from 'element-ui/lib/locale'
locale.use(lang)
const element = {
  install: function (Vue) {
Vue.use(Pagination);
Vue.use(Dialog);
Vue.use(Avatar);
Vue.use(Tag);
// Vue.use(ScorllBar);

Vue.use(Menu);
Vue.use(Submenu);
Vue.use(MenuItem);
Vue.use(MenuItemGroup);
Vue.use(Input);
Vue.use(InputNumber);

Vue.use(Switch);
Vue.use(Select);
Vue.use(Option);
Vue.use(OptionGroup);
Vue.use(Button);
Vue.use(ButtonGroup);
Vue.use(Table);
Vue.use(TableColumn);

Vue.use(Form);
Vue.use(FormItem);
Vue.use(Container);
Vue.use(Header);
Vue.use(Aside);
Vue.use(Main);
Vue.use(Footer);

Vue.use(Link);
Vue.use(Divider);
Vue.use(Image);

Vue.use(Drawer);
Vue.use(Loading.directive);

// Vue.use(Card);
Vue.use(Progress);
Vue.use(Icon);
Vue.use(Row);
Vue.use(Col);
Vue.prototype.$loading = Loading.service;

Vue.prototype.$message = Message;
Vue.use(Radio);
// Vue.use(RadioGroup);
// Vue.use(RadioButton);
Vue.use(Checkbox);
Vue.use(CheckboxButton);
Vue.use(CheckboxGroup);
// Vue.use(Tabs);
// Vue.use(TabPane);
// Vue.use(Tag);
// Vue.use(Tree);
// Vue.use(Alert);
Vue.use(DatePicker);
// Vue.use(TimeSelect);
// Vue.use(TimePicker);
Vue.use(Popover);
Vue.use(Tooltip);
// Vue.use(Breadcrumb);
// Vue.use(BreadcrumbItem);
// Vue.use(Upload);
// Vue.use(Slider);

// Vue.use(Spinner);
// Vue.use(Badge);
// Vue.use(Rate);
// Vue.use(Steps);
// Vue.use(Step);
// Vue.use(Carousel);
// Vue.use(CarouselItem);
// Vue.use(Collapse);
// Vue.use(CollapseItem);
// Vue.use(Cascader);
// Vue.use(ColorPicker);
// Vue.use(Transfer);
// Vue.use(Autocomplete);
// Vue.use(Dropdown);
// Vue.use(DropdownMenu);
// Vue.use(DropdownItem);
// Vue.use(Timeline);
// Vue.use(TimelineItem);
// Vue.use(Calendar);
// Vue.use(Backtop);
Vue.use(PageHeader);
// Vue.use(CascaderPanel);
// Vue.prototype.$msgbox = MessageBox;
// Vue.prototype.$alert = MessageBox.alert;
// Vue.prototype.$confirm = MessageBox.confirm;
// Vue.prototype.$prompt = MessageBox.prompt;
// Vue.prototype.$notify = Notification;

  }
}
export default element