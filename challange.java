package CodingChallange;

import java.util.Scanner;

public class challange {
    public static void main(String[] args) {

        Scanner number = new Scanner(System.in);
        System.out.println("Masukan nilai : ");
        int input_number = number.nextInt(); //memasukan nilai yang diberikan

        System.out.println("Hasil Pangkat tiga dari 1 hingga " + input_number + " adalah: ");

        for (int i = 1; i <= input_number; i++){
            int hasil = (int) Math.pow(i, 3); //menggunakan method math yang berfungsi untuk menghitung dan fungsi pow
            System.out.println("Current Number is : " + i + " and the cube is " + hasil);
        }

    }
}
