package no.uio.ifi.task;

import java.util.concurrent.locks.ReentrantLock;


/* You are allowed to 1. add modifiers to fields and method signatures and 2. add code at the marked places, including removing the following return */
public class LinkedQueue<T> {
    private Node<T> head;
    private Node<T> tail;
    private ReentrantLock lock = new ReentrantLock(true); // make the lock fair

    public int find(T t) {
        lock.lock();
        try {
            Node<T> node = new Node(t);
            while (node != null) {
                if (t == node.content) {
                    return 1; /* return true??? */
                }
                node = node.next;
            }
            return 0;   /* didnt find :( */
        } finally{
            lock.unlock();
        }
    }

    public void insert(T t) {
        lock.lock();
        Node<T> node = new Node(t);
        try {
            if (tail == null) {
                head = node;
            } else {
                tail.next = node;
            }
            tail = node;
        } finally {
            lock.unlock();
        }
    }

    public T delfront() {
        lock.lock();
        try {
            if (head != null) {
                if (head == tail) { tail = null; }
                Node<T> temp = head;
                head = head.next;
                return temp.content;
            }
            return null;
        } finally {
            lock.unlock();
        }
    }
}
