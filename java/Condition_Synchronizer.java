import java.util.concurrent.Semaphore;

public class Condition_Synchronizer {
    private int nr = 0, nw = 0;
    private int dr = 0, dw = 0;
    private Semaphore e = new Semaphore(1);
    private Semaphore r = new Semaphore(5), w = new Semaphore(5);

    public Condition_Synchronizer() {
    }
    
    public void signal() {

    }

    public void reader() {
        while (true) { 
            try {           
                e.acquire();    // P(e)
                if (nw > 0) {
                    dr = dr + 1;
                    e.release();    // V(e)
                    r.acquire();    // P(r)
                }
                nr = nr + 1;
                signal();

                // READ

                e.acquire();
                nr = nr - 1;
                signal();

            } catch (InterruptedException exception) {
                System.out.println(exception);
            }
        }
    }
    
    public static void main(String[] args) throws InterruptedException {
        Condition_Synchronizer obj = new Condition_Synchronizer();
        System.out.println("Final nr: " + obj.nr); 
    }
}