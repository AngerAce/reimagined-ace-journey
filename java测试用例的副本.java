import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

// 商品类
class Product {
    private String name;
    private double price;

    public Product(String name, double price) {
        this.name = name;
        this.price = price;
    }

    public String getName() {
        return name;
    }

    public double getPrice() {
        return price;
    }

    public void setName(String name) {
        this.name = name;
    }

    public void setPrice(double price) {
        this.price = price;
    }
}

// 购物车类
class ShoppingCart {
    private List<Product> products = new ArrayList<>();

    // 添加商品
    public void addProduct(Product product) {
        products.add(product);
    }

    // 删除商品
    public void removeProduct(int index) {
        products.remove(index);
    }

    // 查看购物车
    public void viewShoppingCart() {
        System.out.println("Your Shopping Cart:");
        System.out.println("--------------------");

        double total = 0.0;

        for (int i = 0; i < products.size(); i++) {
            Product product = products.get(i);
            System.out.println((i + 1) + ". " + product.getName() + " - $" + product.getPrice());
            total += product.getPrice();
        }

        System.out.println("--------------------");
        System.out.println("Total: $" + total);
        System.out.println();
    }
}

// 主程序
public class Main {
    public static void main(String[] args) {
        // 创建购物车对象
        ShoppingCart cart = new ShoppingCart();

        // 创建商品对象
        Product product1 = new Product("iPhone X", 999.99);
        Product product2 = new Product("MacBook Pro", 1999.99);
        Product product3 = new Product("Apple Watch", 299.99);

        // 添加商品到购物车
        cart.addProduct(product1);
        cart.addProduct(product2);
        cart.addProduct(product3);

        // 创建一个输入对象
        Scanner scanner = new Scanner(System.in);

        // 循环显示菜单
        while (true) {
            System.out.println("1. Add a product");
            System.out.println("2. Remove a product");
            System.out.println("3. View shopping cart");
            System.out.println("4. Exit");
            System.out.print("Enter your choice: ");

            // 读取用户输入
            int choice = scanner.nextInt();

            switch (choice) {
                case 1:
                    // 添加商品
                    System.out.print("Enter product name: ");
                    String name = scanner.next();
                    System.out.print("Enter product price: ");
                    double price = scanner.nextDouble();
                    Product product = new Product(name, price);
                    cart.addProduct(product);
                    System.out.println("Product added to shopping cart!");
                    break;
                case 2:
                    // 删除商品
                    System.out.print("Enter product index: ");
                    int index = scanner.nextInt() - 1;
                    if (index >= 0 && index < cart.getProducts().size()) {
                        cart.removeProduct(index);
                        System.out.println("Product removed from shopping cart!");
                    } else {
                        System.out.println("Invalid product index!");
                    }
                    break;
                case 3:
                    // 查看购物车
                    cart.viewShoppingCart();
                    break;
                case 4:
                    // 退出程序
